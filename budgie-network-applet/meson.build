project('budgie-network-applet', 'c')

sources = [
    'main.c',
    'applet.c',
    'applet.h',
    'connection_list.h',
    'connection_list.c'
]

gtkdep = dependency('gtk+-3.0')
depends = [
    gtkdep   
]

executable(
    'budgie-network-applet',
    sources,
    dependencies: depends
)