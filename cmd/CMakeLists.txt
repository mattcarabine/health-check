 INCLUDE (FindCouchbaseGo)

  IF (APPLE)
    # On OS X 10.11 (El Capitan) upwards we can no longer use DYLD_LIBRARY_PATH to locate
    # runtime dependancies.
    SET (_ldflags " ${_ldflags} -extldflags '-Wl,-rpath,@executable_path/../lib'")
  ENDIF ()

  GoInstall (TARGET health-check PACKAGE github.com/mattcarabine/health-check/cmd/health
    GOPATH "${PROJECT_SOURCE_DIR}/../../../.." "${GODEPSDIR}"
    LDFLAGS "${_ldflags}"
    INSTALL_PATH bin OUTPUT health-check
    GOVERSION 1.11.4)
