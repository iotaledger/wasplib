@echo off
if not "%1"=="" goto :makeTemplate
echo Please specify smart contract name
goto :xit
:makeTemplate
md %1
xcopy /s helloworld %1
rmdir /s /q %1\pkg
grepl -i -s helloworld %1 %1\*.*
:xit
