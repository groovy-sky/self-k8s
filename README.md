# evo


New-SelfSignedCertificate -Subject 'CN=win10hostvm.westeurope.cloudapp.azure.com' -TextExtension '2.5.29.37={text}1.3.6.1.5.5.7.3.1'

winrm create winrm/config/Listener?Address=*+Transport=HTTPS '@{Hostname="win10hostvm.westeurope.cloudapp.azure.com";CertificateThumbprint="D883CF864F2890A526B98B3547F51E897ED2FEDC"}'