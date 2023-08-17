cmake_minimum_required(VERSION 3.10)
project(restbench_cpp_drogon)

set(CMAKE_MODULE_PATH "${PROJECT_SOURCE_DIR}/build" ${CMAKE_MODULE_PATH})
set(CMAKE_PREFIX_PATH ${CMAKE_BINARY_DIR} ${CMAKE_PREFIX_PATH})

set(CMAKE_BUILD_TYPE Release)
set(EXECUTABLE_OUTPUT_PATH "${PROJECT_SOURCE_DIR}/bin")

set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_TOOLCHAIN_FILE "${PROJECT_SOURCE_DIR}/build/conan_toolchain.cmake")

find_package(Drogon REQUIRED)
find_package(jsoncpp REQUIRED)
find_package(PostgreSQL REQUIRED)

file(GLOB SRC_FILES src/main.cpp src/**/*.h src/**/*.cpp)

add_executable(${PROJECT_NAME} ${SRC_FILES})

target_link_libraries(restbench_cpp_drogon Drogon::Drogon)
