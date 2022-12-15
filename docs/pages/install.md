# Install

Until now, this software does not have installers, so the user can download it, unzip it and place the executable file in the installation location. For some cases adding an environment variable is recommended.

This software is builded to __Mac__ (_amd64_ and _arm64_), __Linux__ (_i386_, _amd64_, and _arm64_), and __Windows__ (_i386_ and _amd64_). All installations files and checksum are disponible at [Last Release Page](https://github.com/gsdenys/pdgen/releases/latest)

Now select the tab related to your operating system to know more about how to instal the [PDGEN](https://gsdenys.github.io/pdgen) on.

<!-- tabs:start -->

#### **Mac**

To install [PDGEN](https://gsdenys.github.io/pdgen) on _Mac_ environment you must to follow the steps below:

?> :pencil2: &nbsp; This tutorial is based in a __M2 (arm64)__ processor. You must do download the best one for your environment.

1. Go to _Download_ home directory and download the version compatible with your environment. There is 2  options for _Mac_, the _arm64_ for _M1_ and _M2_ processors, and _adm64_ for the elders. Alternatively, you can execute the command below to download it. 

```bash
cd ~/Downloads
curl -OL https://github.com/gsdenys/pdgen/releases/download/v1.0.0/pdgen-v1.0.0-darwin-arm64.tar.gz
```

2. Remove any previous [PDGEN](https://gsdenys.github.io/pdgen) installation by deleting the _/usr/local/bin/pdgen_ (if it exists), then extract the archive you just downloaded into _/usr/local/bin_. 

!> You need to have administration permission to execute the next command. 

```bash
sudo rm -f /usr/local/bin/pdgen && sudo tar -C /usr/local/bin -xzf pdgen-v1.0.0-darwin-arm64.tar.gz
```

3. Delete installation file.

```bash
rm pdgen-v1.0.0-darwin-arm64.tar.gz
```

#### **Linux**

To install [PDGEN](https://gsdenys.github.io/pdgen) on _Linux_ environment you must to follow the steps below:

?> :pencil2: &nbsp; This tutorial is based in a __amd64__ processor. You must do download the best one for your environment.

1. Go to _Download_ home directory and download the version compatible with your environment. There is 3  options for _Linux_, the _386_, the _arm64_, and _adm64_. Alternatively, you can execute the command below to download it. 

```bash
cd ~/Downloads
curl -OL https://github.com/gsdenys/pdgen/releases/download/v1.0.0/pdgen-v1.0.0-linux-amd64.tar.gz
```

2. Remove any previous [PDGEN](https://gsdenys.github.io/pdgen) installation by deleting the _/usr/local/bin/pdgen_ (if it exists), then extract the archive you just downloaded into _/usr/local/bin_. 

!> You need to have administration permission to execute the next command. 

```bash
sudo rm -f /usr/local/bin/pdgen && sudo tar -C /usr/local/bin -xzf pdgen-v1.0.0-linux-amd64.tar.gz
```

3. Delete installation file.

```bash
rm pdgen-v1.0.0-linux-amd64.tar.gz
``` 

#### **Windows**

To install [PDGEN](https://gsdenys.github.io/pdgen) on _Windows_ environment you must to follow the steps below:

?> :pencil2: &nbsp; This tutorial is based in a __amd64__ processor. You must do download the best one for your environment.

1. Open Powershell as administrator;

!> The Administrator role is required, other wise it'll not work.

2. Go to _Download_ home directory and download the version compatible with your environment. There is 2 options for _Windows_, the _386_, and _adm64_. Alternatively, you can execute the command below to download it;

```powershell
Start-BitsTransfer -source https://github.com/gsdenys/pdgen/releases/download/v1.0.0/pdgen-v1.0.0-windows-amd64.zip
```

3. Remove any previous [PDGEN](https://gsdenys.github.io/pdgen) installation if it exists. Case it not exists a not important error can be throws;

```powershell
Remove-Item -Path "$env:ProgramFiles\pdgen" -Recurse
```

4. Unzip the downloaded file to the Programs and Files directory;

```powershell
Expand-Archive -Path .\pdgen-v1.0.0-windows-amd64.zip -DestinationPath "$env:ProgramFiles\pdgen"
```

5. Obtains the _Path_ for all users and add the pdgen to the path string;

```powershell
$Path = [Environment]::GetEnvironmentVariable("PATH", "Machine") + [IO.Path]::PathSeparator +  "$env:ProgramFiles\pdgen"
```

6. Update de _Path_ environment variable for all users;

```powershell
[Environment]::SetEnvironmentVariable( "Path", $Path, "Machine" )
```

7. Remove the installation file.

```powershell
Remove-Item -Path .\pdgen-v1.0.0-windows-amd64.zip
```

<!-- tabs:end -->

Once [PDGEN](https://gsdenys.github.io/pdgen) installed, go to the [Usage Page](pages/usage.md) and discovery the power of this tool.