function(nanopb_gen_add)
    # Parse named arguments
    cmake_parse_arguments(
        ARG  # Prefix for parsed variables
        ""   # No boolean flags
        "TARGET;RELPATH;PACKAGE;NANOPB_TAG"  # Single-value arguments
        ""   # No multi-value arguments
        ${ARGN}
    )
    # Validate required arguments
    if(NOT ARG_TARGET)
        message(FATAL_ERROR "nanopb_gen_add: Missing TARGET argument")
    endif()
    if(NOT ARG_RELPATH)
        message(FATAL_ERROR "nanopb_gen_add: Missing RELPATH argument")
    endif()
    if(NOT ARG_PACKAGE)
        message(FATAL_ERROR "nanopb_gen_add: Missing PACKAGE argument")
    endif()
    if(NOT ARG_NANOPB_TAG)
        message(STATUS "nanopb_gen_add: NANOPB_TAG not provided, defaulting to 0.4.9.1")
        set(ARG_NANOPB_TAG "0.4.9.1")
    endif()

    # Convert relative path to an absolute path
    get_filename_component(PROTO_DIR "${CMAKE_CURRENT_SOURCE_DIR}/${ARG_RELPATH}/${ARG_PACKAGE}" ABSOLUTE)

    # Verify that the directory exists
    if(NOT EXISTS ${PROTO_DIR})
        message(FATAL_ERROR "Proto directory ${PROTO_DIR} does not exist.")
    endif()

    # Find all .proto files in the directory (recursively)
    file(GLOB_RECURSE PROTO_FILES "${PROTO_DIR}/*.proto")

    if(NOT PROTO_FILES)
        message(FATAL "No .proto files found in ${PROTO_DIR}")
    endif()

    # Fetch Nanopb if not already included
    include(FetchContent)
    cmake_policy(SET CMP0135 NEW)
    FetchContent_Declare(nanopb URL "https://github.com/nanopb/nanopb/archive/refs/tags/${ARG_NANOPB_TAG}.zip")
    FetchContent_MakeAvailable(nanopb)
    set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${nanopb_SOURCE_DIR}/extra)

    find_package(Python3 REQUIRED)
    find_package(Nanopb REQUIRED)

    nanopb_generate_cpp(
        TARGET ${ARG_TARGET}
        RELPATH ${ARG_RELPATH}
        ${PROTO_FILES}
    )
endfunction()
