TExecFunc arrow_format_type() { return TGenericExec<TPgDirectFunc<&format_type>, false, false, TArgs_format_type_Policy>({}); }
TExecFunc arrow_hashchar() { return TGenericExec<TPgDirectFunc<&hashchar>, true, true, TArgs_hashchar_Policy>({}); }
TExecFunc arrow_hashcharextended() { return TGenericExec<TPgDirectFunc<&hashcharextended>, true, true, TArgs_hashcharextended_Policy>({}); }
TExecFunc arrow_hashenum() { return TGenericExec<TPgDirectFunc<&hashenum>, true, true, TArgs_hashenum_Policy>({}); }
TExecFunc arrow_hashenumextended() { return TGenericExec<TPgDirectFunc<&hashenumextended>, true, true, TArgs_hashenumextended_Policy>({}); }
TExecFunc arrow_hashfloat4() { return TGenericExec<TPgDirectFunc<&hashfloat4>, true, true, TArgs_hashfloat4_Policy>({}); }
TExecFunc arrow_hashfloat4extended() { return TGenericExec<TPgDirectFunc<&hashfloat4extended>, true, true, TArgs_hashfloat4extended_Policy>({}); }
TExecFunc arrow_hashfloat8() { return TGenericExec<TPgDirectFunc<&hashfloat8>, true, true, TArgs_hashfloat8_Policy>({}); }
TExecFunc arrow_hashfloat8extended() { return TGenericExec<TPgDirectFunc<&hashfloat8extended>, true, true, TArgs_hashfloat8extended_Policy>({}); }
TExecFunc arrow_hashint2() { return TGenericExec<TPgDirectFunc<&hashint2>, true, true, TArgs_hashint2_Policy>({}); }
TExecFunc arrow_hashint2extended() { return TGenericExec<TPgDirectFunc<&hashint2extended>, true, true, TArgs_hashint2extended_Policy>({}); }
TExecFunc arrow_hashint4() { return TGenericExec<TPgDirectFunc<&hashint4>, true, true, TArgs_hashint4_Policy>({}); }
TExecFunc arrow_hashint4extended() { return TGenericExec<TPgDirectFunc<&hashint4extended>, true, true, TArgs_hashint4extended_Policy>({}); }
TExecFunc arrow_hashint8() { return TGenericExec<TPgDirectFunc<&hashint8>, true, true, TArgs_hashint8_Policy>({}); }
TExecFunc arrow_hashint8extended() { return TGenericExec<TPgDirectFunc<&hashint8extended>, true, true, TArgs_hashint8extended_Policy>({}); }
TExecFunc arrow_hashname() { return TGenericExec<TPgDirectFunc<&hashname>, true, true, TArgs_hashname_Policy>({}); }
TExecFunc arrow_hashnameextended() { return TGenericExec<TPgDirectFunc<&hashnameextended>, true, true, TArgs_hashnameextended_Policy>({}); }
TExecFunc arrow_hashoid() { return TGenericExec<TPgDirectFunc<&hashoid>, true, true, TArgs_hashoid_Policy>({}); }
TExecFunc arrow_hashoidextended() { return TGenericExec<TPgDirectFunc<&hashoidextended>, true, true, TArgs_hashoidextended_Policy>({}); }
TExecFunc arrow_hashoidvector() { return TGenericExec<TPgDirectFunc<&hashoidvector>, true, true, TArgs_hashoidvector_Policy>({}); }
TExecFunc arrow_hashoidvectorextended() { return TGenericExec<TPgDirectFunc<&hashoidvectorextended>, true, true, TArgs_hashoidvectorextended_Policy>({}); }
TExecFunc arrow_hashtext() { return TGenericExec<TPgDirectFunc<&hashtext>, true, true, TArgs_hashtext_Policy>({}); }
TExecFunc arrow_hashtextextended() { return TGenericExec<TPgDirectFunc<&hashtextextended>, true, true, TArgs_hashtextextended_Policy>({}); }
TExecFunc arrow_hashvarlena() { return TGenericExec<TPgDirectFunc<&hashvarlena>, true, true, TArgs_hashvarlena_Policy>({}); }
TExecFunc arrow_hashvarlenaextended() { return TGenericExec<TPgDirectFunc<&hashvarlenaextended>, true, true, TArgs_hashvarlenaextended_Policy>({}); }
TExecFunc arrow_jsonb_agg_finalfn() { return TGenericExec<TPgDirectFunc<&jsonb_agg_finalfn>, false, false, TArgs_jsonb_agg_finalfn_Policy>({}); }
TExecFunc arrow_jsonb_agg_transfn() { return TGenericExec<TPgDirectFunc<&jsonb_agg_transfn>, false, true, TArgs_jsonb_agg_transfn_Policy>({}); }
TExecFunc arrow_jsonb_bool() { return TGenericExec<TPgDirectFunc<&jsonb_bool>, true, true, TArgs_jsonb_bool_Policy>({}); }
TExecFunc arrow_jsonb_build_array_noargs() { return TGenericExec<TPgDirectFunc<&jsonb_build_array_noargs>, false, false, TArgs_jsonb_build_array_noargs_Policy>({}); }
TExecFunc arrow_jsonb_build_object_noargs() { return TGenericExec<TPgDirectFunc<&jsonb_build_object_noargs>, false, false, TArgs_jsonb_build_object_noargs_Policy>({}); }
TExecFunc arrow_jsonb_float4() { return TGenericExec<TPgDirectFunc<&jsonb_float4>, true, true, TArgs_jsonb_float4_Policy>({}); }
TExecFunc arrow_jsonb_float8() { return TGenericExec<TPgDirectFunc<&jsonb_float8>, true, true, TArgs_jsonb_float8_Policy>({}); }
TExecFunc arrow_jsonb_in() { return TGenericExec<TPgDirectFunc<&jsonb_in>, true, false, TArgs_jsonb_in_Policy>({}); }
TExecFunc arrow_jsonb_int2() { return TGenericExec<TPgDirectFunc<&jsonb_int2>, true, true, TArgs_jsonb_int2_Policy>({}); }
TExecFunc arrow_jsonb_int4() { return TGenericExec<TPgDirectFunc<&jsonb_int4>, true, true, TArgs_jsonb_int4_Policy>({}); }
TExecFunc arrow_jsonb_int8() { return TGenericExec<TPgDirectFunc<&jsonb_int8>, true, true, TArgs_jsonb_int8_Policy>({}); }
TExecFunc arrow_jsonb_numeric() { return TGenericExec<TPgDirectFunc<&jsonb_numeric>, true, false, TArgs_jsonb_numeric_Policy>({}); }
TExecFunc arrow_jsonb_object() { return TGenericExec<TPgDirectFunc<&jsonb_object>, true, false, TArgs_jsonb_object_Policy>({}); }
TExecFunc arrow_jsonb_object_agg_finalfn() { return TGenericExec<TPgDirectFunc<&jsonb_object_agg_finalfn>, false, false, TArgs_jsonb_object_agg_finalfn_Policy>({}); }
TExecFunc arrow_jsonb_object_agg_transfn() { return TGenericExec<TPgDirectFunc<&jsonb_object_agg_transfn>, false, true, TArgs_jsonb_object_agg_transfn_Policy>({}); }
TExecFunc arrow_jsonb_object_two_arg() { return TGenericExec<TPgDirectFunc<&jsonb_object_two_arg>, true, false, TArgs_jsonb_object_two_arg_Policy>({}); }
TExecFunc arrow_jsonb_out() { return TGenericExec<TPgDirectFunc<&jsonb_out>, true, false, TArgs_jsonb_out_Policy>({}); }
TExecFunc arrow_jsonb_recv() { return TGenericExec<TPgDirectFunc<&jsonb_recv>, true, false, TArgs_jsonb_recv_Policy>({}); }
TExecFunc arrow_jsonb_send() { return TGenericExec<TPgDirectFunc<&jsonb_send>, true, false, TArgs_jsonb_send_Policy>({}); }
TExecFunc arrow_jsonb_subscript_handler() { return TGenericExec<TPgDirectFunc<&jsonb_subscript_handler>, true, true, TArgs_jsonb_subscript_handler_Policy>({}); }
TExecFunc arrow_jsonb_typeof() { return TGenericExec<TPgDirectFunc<&jsonb_typeof>, true, false, TArgs_jsonb_typeof_Policy>({}); }
TExecFunc arrow_oideq() { return TGenericExec<TPgDirectFunc<&oideq>, true, true, TArgs_oideq_Policy>({}); }
TExecFunc arrow_oidge() { return TGenericExec<TPgDirectFunc<&oidge>, true, true, TArgs_oidge_Policy>({}); }
TExecFunc arrow_oidgt() { return TGenericExec<TPgDirectFunc<&oidgt>, true, true, TArgs_oidgt_Policy>({}); }
TExecFunc arrow_oidin() { return TGenericExec<TPgDirectFunc<&oidin>, true, true, TArgs_oidin_Policy>({}); }
TExecFunc arrow_oidlarger() { return TGenericExec<TPgDirectFunc<&oidlarger>, true, true, TArgs_oidlarger_Policy>({}); }
TExecFunc arrow_oidle() { return TGenericExec<TPgDirectFunc<&oidle>, true, true, TArgs_oidle_Policy>({}); }
TExecFunc arrow_oidlt() { return TGenericExec<TPgDirectFunc<&oidlt>, true, true, TArgs_oidlt_Policy>({}); }
TExecFunc arrow_oidne() { return TGenericExec<TPgDirectFunc<&oidne>, true, true, TArgs_oidne_Policy>({}); }
TExecFunc arrow_oidout() { return TGenericExec<TPgDirectFunc<&oidout>, true, false, TArgs_oidout_Policy>({}); }
TExecFunc arrow_oidrecv() { return TGenericExec<TPgDirectFunc<&oidrecv>, true, true, TArgs_oidrecv_Policy>({}); }
TExecFunc arrow_oidsend() { return TGenericExec<TPgDirectFunc<&oidsend>, true, false, TArgs_oidsend_Policy>({}); }
TExecFunc arrow_oidsmaller() { return TGenericExec<TPgDirectFunc<&oidsmaller>, true, true, TArgs_oidsmaller_Policy>({}); }
TExecFunc arrow_oidvectoreq() { return TGenericExec<TPgDirectFunc<&oidvectoreq>, true, true, TArgs_oidvectoreq_Policy>({}); }
TExecFunc arrow_oidvectorge() { return TGenericExec<TPgDirectFunc<&oidvectorge>, true, true, TArgs_oidvectorge_Policy>({}); }
TExecFunc arrow_oidvectorgt() { return TGenericExec<TPgDirectFunc<&oidvectorgt>, true, true, TArgs_oidvectorgt_Policy>({}); }
TExecFunc arrow_oidvectorin() { return TGenericExec<TPgDirectFunc<&oidvectorin>, true, false, TArgs_oidvectorin_Policy>({}); }
TExecFunc arrow_oidvectorle() { return TGenericExec<TPgDirectFunc<&oidvectorle>, true, true, TArgs_oidvectorle_Policy>({}); }
TExecFunc arrow_oidvectorlt() { return TGenericExec<TPgDirectFunc<&oidvectorlt>, true, true, TArgs_oidvectorlt_Policy>({}); }
TExecFunc arrow_oidvectorne() { return TGenericExec<TPgDirectFunc<&oidvectorne>, true, true, TArgs_oidvectorne_Policy>({}); }
TExecFunc arrow_oidvectorout() { return TGenericExec<TPgDirectFunc<&oidvectorout>, true, false, TArgs_oidvectorout_Policy>({}); }
TExecFunc arrow_oidvectorrecv() { return TGenericExec<TPgDirectFunc<&oidvectorrecv>, true, false, TArgs_oidvectorrecv_Policy>({}); }
TExecFunc arrow_oidvectorsend() { return TGenericExec<TPgDirectFunc<&oidvectorsend>, true, false, TArgs_oidvectorsend_Policy>({}); }
TExecFunc arrow_oidvectortypes() { return TGenericExec<TPgDirectFunc<&oidvectortypes>, true, false, TArgs_oidvectortypes_Policy>({}); }
TExecFunc arrow_pgsql_version() { return TGenericExec<TPgDirectFunc<&pgsql_version>, true, false, TArgs_pgsql_version_Policy>({}); }
TExecFunc arrow_to_jsonb() { return TGenericExec<TPgDirectFunc<&to_jsonb>, true, false, TArgs_to_jsonb_Policy>({}); }
