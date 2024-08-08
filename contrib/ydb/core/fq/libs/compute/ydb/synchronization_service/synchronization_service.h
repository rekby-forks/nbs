#pragma once

#include <contrib/ydb/core/fq/libs/config/protos/compute.pb.h>

#include <contrib/ydb/library/security/ydb_credentials_provider_factory.h>
#include <contrib/ydb/library/grpc/actor_client/grpc_service_settings.h>
#include <contrib/ydb/core/fq/libs/config/protos/common.pb.h>
#include <contrib/ydb/core/fq/libs/signer/signer.h>
#include <contrib/ydb/core/fq/libs/shared_resources/shared_resources.h>

#include <contrib/ydb/library/actors/core/actor.h>

namespace NFq {

std::unique_ptr<NActors::IActor> CreateSynchronizationServiceActor(const NConfig::TCommonConfig& commonConfig,
                                                                   const NConfig::TComputeConfig& computeConfig,
                                                                   const TSigner::TPtr& signer,
                                                                   const TYqSharedResources::TPtr& yqSharedResources,
                                                                   const NKikimr::TYdbCredentialsProviderFactory& credentialsProviderFactory,
                                                                   const ::NMonitoring::TDynamicCounterPtr& counters);

}
