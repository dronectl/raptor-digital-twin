cmake_minimum_required(VERSION 3.22)

message(STATUS "Configuring config.h header")
configure_file(${CMAKE_CURRENT_SOURCE_DIR}/config.h.in config.h @ONLY)
include_directories(${CMAKE_CURRENT_BINARY_DIR})

add_executable(${PROJECT_NAME} main.c)

target_compile_options(
  ${PROJECT_NAME}
  PUBLIC -Wno-unused-parameter -Wpedantic -fno-builtin -Wall -Wextra -ffunction-sections -fdata-sections -fomit-frame-pointer
  PUBLIC $<$<CONFIG:Debug>:-DDEBUG>
)

set_target_properties(${PROJECT_NAME} PROPERTIES RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}")
