@echo off
if not "%1"=="" goto :makeTemplate
echo Please specify smart contract name
goto :xit
:makeTemplate
copy helloworld.go %1.go
md %1
copy helloworld\helloworld.go %1\%1.go
grepl -i -s helloworld %1 %1.go
:xit
