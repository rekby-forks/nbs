#pragma once

#include "public.h"

#include <cloud/blockstore/public/api/protos/encryption.pb.h>
#include <cloud/blockstore/public/api/protos/volume.pb.h>

#include <cloud/blockstore/libs/diagnostics/public.h>
#include <cloud/blockstore/libs/service/public.h>
#include <cloud/storage/core/libs/common/error.h>

namespace NCloud::NBlockStore {

////////////////////////////////////////////////////////////////////////////////

IBlockStorePtr CreateEncryptionClient(
    IBlockStorePtr client,
    ILoggingServicePtr logging,
    IEncryptorPtr encryptor,
    NProto::TEncryptionDesc encryptionDesc);

IBlockStorePtr CreateSnapshotEncryptionClient(
    IBlockStorePtr client,
    ILoggingServicePtr logging,
    NProto::TEncryptionDesc encryptionDesc);

////////////////////////////////////////////////////////////////////////////////

struct IEncryptionClientFactory
{
    using TResponse = TResultOrError<IBlockStorePtr>;

    virtual ~IEncryptionClientFactory() = default;

    virtual NThreading::TFuture<TResponse> CreateEncryptionClient(
        IBlockStorePtr client,
        const NProto::TEncryptionSpec& encryptionSpec,
        const TString& diskId) = 0;
};

////////////////////////////////////////////////////////////////////////////////

IEncryptionClientFactoryPtr CreateEncryptionClientFactory(
    ILoggingServicePtr logging,
    IEncryptionKeyProviderPtr encryptionKeyProvider);

}   // namespace NCloud::NBlockStore
