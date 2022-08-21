# Bitdefender GravityZone Import/Export API
In order to import policies from one instance of Bitdefender GravityZone appliance to another one, this app will use GZ API to export the required policies from the GZ appliance and will import it to new instance.


### HOW TO
1. Make sure you have created an API key on both Instance of GZ Appliances. [Read Official Documentation](https://www.bitdefender.com/business/support/en/77209-125280-getting-started.html#UUID-e6befdd4-3eb1-4b6e-cc6c-19bdd16847b4_section-idm4640169987334432655171029621) for creating API Key
2. Download Policy Migration tools:
   - **Linux**:
     - [i386 architecture](https://github.com/javadmohebbi/gobdgz/raw/master/dist/linux/x64/policy-migration)
     - [amd64 architecture](https://github.com/javadmohebbi/gobdgz/raw/master/dist/linux/x64/policy-migration)
   - **Windows**:
     - [Windows 32bit](https://github.com/javadmohebbi/gobdgz/raw/master/dist/windows/x86/policy-migration.exe)
     - [Windows 64bit](https://github.com/javadmohebbi/gobdgz/raw/master/dist/windows/x64/policy-migration.exe)
   - **macOS**: [amd64 architecture](https://github.com/javadmohebbi/gobdgz/raw/master/dist/macOS/x64/policy-migration)
3. Create a file called `config.json` and use the below template and place your own here:
```
{
    "SRC": {
        "SERVER": "192.168.1.1",
        "API_KEY": "928eb517ef0cc0f84c910c38f14f1bb4b5fac87af099e704b40d71c1cf61ac24",
        "POLICIES": [
            "Policy1",
            "Policy2",
            "Policy3"
        ]
    },
    "DST": {
        "SERVER": "192.168.1.2",
        "API_KEY": "d6c812141b354aed5a04cc69cbde132385c05437c17a6e33bb8dcf7060d393d0"
    }
}
```
4. Run the following command to migrate from **SRC** server to **DST** server.
    - ON Linux & macOS ```chmod +x -v policy-migration```
    - Linux: ```$ policy-migration -config /path/to/config.json```
    - Windows: ```C:\ > policy-migration.exe -config \path\to\config.json```
    - macOS: ```$ policy-migration -config /path/to/config.json```

5- If all the requirements are met, policies would be migrated from old server to new one.