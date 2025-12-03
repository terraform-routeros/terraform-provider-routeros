#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/system/leds get [print show-ids]]
terraform import routeros_system_led.sfp1 '*1'
#Or you can import a resource using one of its attributes
terraform import routeros_system_led.sfp1 "name=xxx"
