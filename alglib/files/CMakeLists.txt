# Set  the minimum  required version  of cmake  for a  project.
cmake_minimum_required(VERSION 2.6)

set(ALGLIB_VERSION "" CACHE STRING "The package version")

file(GLOB_RECURSE HDR RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} src/*.h)
file(GLOB_RECURSE SRC RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} src/*.cpp)

add_library(alglib SHARED ${SRC})
set_target_properties(alglib PROPERTIES SUFFIX "-${ALGLIB_VERSION}${CMAKE_SHARED_LIBRARY_SUFFIX}")

add_executable(test_c tests/test_c.cpp)
add_executable(test_i tests/test_i.cpp)
# add_executable(test_x tests/test_x.cpp)
set_target_properties(test_c PROPERTIES COMPILE_FLAGS "-DAE_USE_ALLOC_COUNTER -DAE_DEBUG4POSIX")
set_target_properties(test_i PROPERTIES COMPILE_FLAGS "-DAE_USE_ALLOC_COUNTER -DAE_DEBUG4POSIX")
# set_target_properties(test_x PROPERTIES COMPILE_FLAGS "-DAE_USE_ALLOC_COUNTER -DAE_DEBUG4POSIX")
include_directories(src/)
target_link_libraries(test_c alglib)
target_link_libraries(test_i alglib)
# target_link_libraries(test_x alglib)

install(TARGETS alglib LIBRARY DESTINATION lib${LIB_SUFFIX})
install(FILES ${HDR} DESTINATION include/alglib)
