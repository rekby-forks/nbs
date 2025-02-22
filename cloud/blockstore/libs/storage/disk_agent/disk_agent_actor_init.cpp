#include "disk_agent_actor.h"

#include <cloud/blockstore/libs/diagnostics/request_stats.h>
#include <cloud/storage/core/libs/common/error.h>
#include <cloud/storage/core/libs/common/format.h>
#include <cloud/storage/core/libs/diagnostics/public.h>

#include <util/generic/vector.h>
#include <util/stream/str.h>
#include <util/string/builder.h>
#include <util/string/cast.h>

namespace NCloud::NBlockStore::NStorage {

using namespace NActors;

////////////////////////////////////////////////////////////////////////////////

void TDiskAgentActor::InitAgent(const TActorContext& ctx)
{
    State = std::make_unique<TDiskAgentState>(
        Config,
        AgentConfig,
        RdmaConfig,
        Spdk,
        Allocator,
        StorageProvider,
        ProfileLog,
        BlockDigestGenerator,
        Logging,
        RdmaServer,
        NvmeManager);

    Y_DEBUG_ABORT_UNLESS(
        OldRequestCounters.Delayed && OldRequestCounters.Rejected &&
        OldRequestCounters.Already);
    TRdmaTargetConfig rdmaTargetConfig{
        RejectLateRequestsAtDiskAgentEnabled,
        OldRequestCounters};

    auto result = State->Initialize(std::move(rdmaTargetConfig));

    auto* actorSystem = ctx.ActorSystem();
    auto replyTo = ctx.SelfID;

    result.Subscribe([=] (auto future) {
        using TCompletionEvent = TEvDiskAgentPrivate::TEvInitAgentCompleted;

        NProto::TError error;

        try {
            TDiskAgentState::TInitializeResult r = future.ExtractValue();

            auto response = std::make_unique<TCompletionEvent>(
                std::move(r.Configs),
                std::move(r.Errors),
                std::move(r.ConfigMismatchErrors));

            actorSystem->Send(
                new IEventHandle(
                    replyTo,
                    replyTo,
                    response.release()));
        } catch (const TServiceError& e) {
            error = MakeError(e.GetCode(), TString(e.GetMessage()));
        } catch (...) {
            error = MakeError(E_FAIL, CurrentExceptionMessage());
        }

        if (error.GetCode()) {
            auto response = std::make_unique<TCompletionEvent>(error);

            actorSystem->Send(
                new IEventHandle(
                    replyTo,
                    replyTo,
                    response.release()));
        }
    });
}

void TDiskAgentActor::HandleInitAgentCompleted(
    const TEvDiskAgentPrivate::TEvInitAgentCompleted::TPtr& ev,
    const TActorContext& ctx)
{
    auto* msg = ev->Get();

    for (const auto& error: msg->Errors) {
        LOG_WARN_S(ctx, TBlockStoreComponents::DISK_AGENT, error);
    }

    // Crit events that reported on startup have issue with them being invisible
    // on second restart. Here, we schedule the event to allow monitoring
    // initially to read counters without event and then with the event.
    for (const auto& configMismatchError: msg->ConfigMismatchErrors) {
        const TDuration startupCritEventDelay = UpdateCountersInterval * 2;
        ctx.Schedule(
            startupCritEventDelay,
            new TEvDiskAgentPrivate::TEvReportDelayedDiskAgentConfigMismatch(
                configMismatchError));
    }

    if (const auto& error = msg->GetError(); HasError(error)) {
        LOG_ERROR_S(ctx, TBlockStoreComponents::DISK_AGENT,
            "DiskAgent initialization failed. Error: " << FormatError(error).data());
    } else {
        TStringStream out;
        for (const auto& config: msg->Configs) {
            out << config.GetDeviceName()
                << "(" << FormatByteSize(config.GetBlocksCount() * config.GetBlockSize())
                << "); ";
        }

        LOG_INFO_S(ctx, TBlockStoreComponents::DISK_AGENT,
            "Initialization completed. Devices found: " << out.Str());
    }

    // resend pending requests
    SendPendingRequests(ctx, PendingRequests);

    if (msg->Configs.empty()) {
        LOG_INFO(
            ctx,
            TBlockStoreComponents::DISK_AGENT,
            "No devices: become idle");

        Become(&TThis::StateIdle);

        return;
    }

    Become(&TThis::StateWork);

    NCloud::Send(
        ctx,
        MakeDiskRegistryProxyServiceId(),
        std::make_unique<TEvDiskRegistryProxy::TEvSubscribeRequest>(
            ctx.SelfID));

    ScheduleUpdateStats(ctx);

    RunSessionCacheActor(ctx);
}

}   // namespace NCloud::NBlockStore::NStorage
