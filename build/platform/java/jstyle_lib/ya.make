RESOURCES_LIBRARY()

OWNER(heretic)

IF (USE_SYSTEM_JSTYLE_LIB)
    MESSAGE(WARNING System java codestyle library $USE_SYSTEM_JSTYLE_LIB will be used)
ELSE()
    DECLARE_EXTERNAL_RESOURCE(JSTYLE_LIB sbr:4978388684)
ENDIF()

END()