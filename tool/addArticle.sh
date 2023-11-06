#!/bin/bash
echo "输入归属目录及文件名"
read path
cd hugo
hugo new post/$path
