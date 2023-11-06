#!/bin/bash
rm -rf config/*
mv web/protect .
rm -rf web/*
rm -rf hugo/assets/
rm -rf hugo/resources/
rm -rf hugo/public/*
mv protect web/
