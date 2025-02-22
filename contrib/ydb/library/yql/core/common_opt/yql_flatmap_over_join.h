#pragma once

#include <contrib/ydb/library/yql/core/expr_nodes/yql_expr_nodes.h>

#include <contrib/ydb/library/yql/ast/yql_expr.h>

#include <util/generic/strbuf.h>


namespace NYql {

NNodes::TExprBase FlatMapOverEquiJoin(const NNodes::TCoFlatMapBase& node, TExprContext& ctx, const TParentsMap& parentsMap, bool multiUsage);

} // NYql
