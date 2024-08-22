#pragma once

#include "public.h"

#include <cloud/storage/core/libs/kikimr/public.h>

#include <cloud/storage/core/protos/authorization_mode.pb.h>

namespace NCloud::NStorage {

////////////////////////////////////////////////////////////////////////////////

NActors::IActorPtr CreateAuthorizerActor(
    int component,
    TString counterId,
    TString folderId,
    TString containerId,
    NProto::EAuthorizationMode authMode,
    bool checkAuthorization);

}   // namespace NCloud::NStorage

