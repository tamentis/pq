package pq

// Generated via massaging this catalog query:
//
//     SELECT 'T_' || typname || ' = ' || oid
//     FROM pg_type WHERE oid < 10000
//     ORDER BY oid;
//
// This should probably be done one per release.  Postgres does not
// re-appropriate the system OID space below 10000 as a general rule.

type Oid uint32

const (
	T_bool             = 16
	T_bytea            = 17
	T_char             = 18
	T_name             = 19
	T_int8             = 20
	T_int2             = 21
	T_int2vector       = 22
	T_int4             = 23
	T_regproc          = 24
	T_text             = 25
	T_oid              = 26
	T_tid              = 27
	T_xid              = 28
	T_cid              = 29
	T_oidvector        = 30
	T_pg_type          = 71
	T_pg_attribute     = 75
	T_pg_proc          = 81
	T_pg_class         = 83
	T_json             = 114
	T_xml              = 142
	T__xml             = 143
	T_pg_node_tree     = 194
	T__json            = 199
	T_smgr             = 210
	T_point            = 600
	T_lseg             = 601
	T_path             = 602
	T_box              = 603
	T_polygon          = 604
	T_line             = 628
	T__line            = 629
	T_cidr             = 650
	T__cidr            = 651
	T_float4           = 700
	T_float8           = 701
	T_abstime          = 702
	T_reltime          = 703
	T_tinterval        = 704
	T_unknown          = 705
	T_circle           = 718
	T__circle          = 719
	T_money            = 790
	T__money           = 791
	T_macaddr          = 829
	T_inet             = 869
	T__bool            = 1000
	T__bytea           = 1001
	T__char            = 1002
	T__name            = 1003
	T__int2            = 1005
	T__int2vector      = 1006
	T__int4            = 1007
	T__regproc         = 1008
	T__text            = 1009
	T__tid             = 1010
	T__xid             = 1011
	T__cid             = 1012
	T__oidvector       = 1013
	T__bpchar          = 1014
	T__varchar         = 1015
	T__int8            = 1016
	T__point           = 1017
	T__lseg            = 1018
	T__path            = 1019
	T__box             = 1020
	T__float4          = 1021
	T__float8          = 1022
	T__abstime         = 1023
	T__reltime         = 1024
	T__tinterval       = 1025
	T__polygon         = 1027
	T__oid             = 1028
	T_aclitem          = 1033
	T__aclitem         = 1034
	T__macaddr         = 1040
	T__inet            = 1041
	T_bpchar           = 1042
	T_varchar          = 1043
	T_date             = 1082
	T_time             = 1083
	T_timestamp        = 1114
	T__timestamp       = 1115
	T__date            = 1182
	T__time            = 1183
	T_timestamptz      = 1184
	T__timestamptz     = 1185
	T_interval         = 1186
	T__interval        = 1187
	T__numeric         = 1231
	T_pg_database      = 1248
	T__cstring         = 1263
	T_timetz           = 1266
	T__timetz          = 1270
	T_bit              = 1560
	T__bit             = 1561
	T_varbit           = 1562
	T__varbit          = 1563
	T_numeric          = 1700
	T_refcursor        = 1790
	T__refcursor       = 2201
	T_regprocedure     = 2202
	T_regoper          = 2203
	T_regoperator      = 2204
	T_regclass         = 2205
	T_regtype          = 2206
	T__regprocedure    = 2207
	T__regoper         = 2208
	T__regoperator     = 2209
	T__regclass        = 2210
	T__regtype         = 2211
	T_record           = 2249
	T_cstring          = 2275
	T_any              = 2276
	T_anyarray         = 2277
	T_void             = 2278
	T_trigger          = 2279
	T_language_handler = 2280
	T_internal         = 2281
	T_opaque           = 2282
	T_anyelement       = 2283
	T__record          = 2287
	T_anynonarray      = 2776
	T_pg_authid        = 2842
	T_pg_auth_members  = 2843
	T__txid_snapshot   = 2949
	T_uuid             = 2950
	T__uuid            = 2951
	T_txid_snapshot    = 2970
	T_fdw_handler      = 3115
	T_anyenum          = 3500
	T_tsvector         = 3614
	T_tsquery          = 3615
	T_gtsvector        = 3642
	T__tsvector        = 3643
	T__gtsvector       = 3644
	T__tsquery         = 3645
	T_regconfig        = 3734
	T__regconfig       = 3735
	T_regdictionary    = 3769
	T__regdictionary   = 3770
	T_anyrange         = 3831
	T_int4range        = 3904
	T__int4range       = 3905
	T_numrange         = 3906
	T__numrange        = 3907
	T_tsrange          = 3908
	T__tsrange         = 3909
	T_tstzrange        = 3910
	T__tstzrange       = 3911
	T_daterange        = 3912
	T__daterange       = 3913
	T_int8range        = 3926
	T__int8range       = 3927
)
