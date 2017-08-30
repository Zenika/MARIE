util = require "util"

SSID = 'Zenika'
PASSWORD = 'ZestParti!'

actions = { {name="on"}, {name="off"} }
getters = { }

local pin = 7
local on = false

gpio.mode(pin, gpio.OUTPUT)
gpio.write(pin, gpio.LOW)

function init_mqtt()
    m = mqtt.Client("thing", 120)
    m:on("connect", function(client) print("mqtt connected") end)
    m:on("message", function(client, topic, data)
        if isAction(topic, "on") then
            gpio.write(pin, gpio.HIGH)
        else
            gpio.write(pin, gpio.LOW)
        end
        print(topic .. ":")
        if data ~= nil then
            print(data)
        end
    end)
    
    m:connect("10.0.10.3", 1883, 0, function(client)
        register(client, "Light", "light", "couloir", actions, getters)
    end)
end

function wait_for_wifi_conn()
    tmr.alarm(1, 1000, 1, function ()
        if wifi.sta.getip() == nil then
            print("Waiting for wifi connection")
        else
            print("Connected")
            tmr.stop(1)
            init_mqtt()
        end
    end)
end

function connect()
    wifi.setmode(wifi.STATION)
    wifi.sta.config(SSID, PASSWORD, 1)
    wait_for_wifi_conn()
end
connect()
