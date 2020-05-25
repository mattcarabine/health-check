if ("${PRODUCT_VERSION}" VERSION_LESS "6.5.0" AND
    NOT "${PRODUCT_VERSION}" VERSION_EQUAL "0.0.0")
  set (PROXY_STRATEGY "local")
else()
  set (PROXY_STRATEGY "sticky")
endif()

# for actual install
set (BIN_PREFIX "")
configure_file (pluggable-ui-health.json.in pluggable-ui-health.json)
install(FILES "${PROJECT_BINARY_DIR}/ui/pluggable-ui-health.json"
        DESTINATION etc/couchbase)

# for cluster_run
set (BIN_PREFIX "${CMAKE_CURRENT_SOURCE_DIR}/")
configure_file (pluggable-ui-health.json.in pluggable-ui-health.cluster_run.json)
file(COPY "${PROJECT_BINARY_DIR}/ui/pluggable-ui-health.cluster_run.json"
        DESTINATION  "${PROJECT_BINARY_DIR}/../../../../../cluster_run_ui_plugins")

# install all files including third party
install(DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/health-ui
       DESTINATION lib)