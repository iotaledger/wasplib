@echo off
for /d %%f in (*.) do call build_go.bat %%f %1
