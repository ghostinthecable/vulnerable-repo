#!/usr/bin/env python3
import os
import pickle

# VULNERABLE: eval on user input, dangerous pickle loading

cmd = input("Enter a command to run: ")
print(eval(cmd))      # code execution vulnerability

with open("payload.pkl", "rb") as f:
    data = pickle.load(f)   # arbitrary code execution via pickle

print("Loaded:", data)
