#pragma once

#include "command.h"

namespace NCloud::NBlockStore::NClient {

////////////////////////////////////////////////////////////////////////////////

TCommandPtr NewCreateCheckpointCommand(IBlockStorePtr client);

}   // namespace NCloud::NBlockStore::NClient