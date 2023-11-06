#!/bin/bash
# Ca文件夹未创建
tar czvf akBlogInfo.tar.gz hugo/themes/hugo-theme-stack/assets/img/avatar.png  config/cert/ hugo/content/post/ 
rm -rf hugo/themes/hugo-theme-stack/assets/img/avatar.png  config/cert/
