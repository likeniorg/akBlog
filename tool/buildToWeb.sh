#!/bin/bash
mv web/protect .
rm -rf web/*
cd hugo
hugo
cp -r ./public/* ../web/
