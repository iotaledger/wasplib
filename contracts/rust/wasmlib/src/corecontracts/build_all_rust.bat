@echo off
for /d %%f in (*.) do call build_rust.bat %%f %1
