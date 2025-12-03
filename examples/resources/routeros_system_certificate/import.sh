#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/certificate get [print show-ids]]
#If you plan to manipulate the certificate requiring signing, you need to correctly fill in the sign{} section.
#Changes in the sign{} section will not cause changes in the certificate. It's not a bug, it's a feature!
terraform import routeros_system_certificate.client *9D
#Or you can import a resource using one of its attributes
terraform import routeros_system_certificate.client "name=xxx"
