@echo off
if not "%1"=="" goto :makeTemplate
echo Please specify smart contract name
goto :xit
:makeTemplate
copy HelloWorld.java %1.java
grepl -i HelloWorld %1 %1.java
:xit
