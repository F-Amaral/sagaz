#!/bin/bash

## Linux install procedure 
## Ubuntu
#
function add_libvips_repos_ubuntu (){
    echo "Adding libvips repositories"
    sudo add-apt-repository -uy ppa:strukturag/libde265 &&\
    sudo add-apt-repository -uy ppa:strukturag/libheif &&\
    sudo add-apt-repository -uy ppa:tonimelisma/ppa &&\
    return ?
}

function install_libvips_ubuntu (){
    echo "Installing libvips"
    sudo apt-get install -y libvips-dev 
    return ?
}

function ubuntu_install_deps () {

    # Check if libvips-dev is available to install
    local is_avail="$(apt -qq list libvips-dev 2> /dev/null)"
    if [ -n "$is_avail" ] ; then
        # Pacote presente no repo 
        if echo "$is_avail" | grep '[installed]' -q; then
            echo "libvips-dev is already installed"
            return 0
        else 
            echo "libvips-dev is not installed"
            install_libvips_ubuntu
        fi
    else
        echo "libvips-dev is not available"	
        add_libvips_repos_ubuntu && install_libvips_ubuntu
        return ?
    fi

    return 0
}


# check if  operating system is Ubuntu Linux or macOS
# 
OSTYPE="$(uname)"

if [ "$OSTYPE" == "Darwin" ];
then
    if [ -x "$(command -v brew)" ]; then
        brew install vips pkg-config
    else 
        echo "brew not installed"
        exit
    fi

elif [ "$OSTYPE" == "Linux" ];
then
   
    if [ "$(lsb_release -si)" == "Ubuntu" ];
    then
        ubuntu_install_deps
    else
        echo "This script is only for Ubuntu"
        exit 1
    fi
else
    echo "Operating system not supported"
    exit 1
fi
