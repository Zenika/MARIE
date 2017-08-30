function doSubscribeAction(client, thingType, location, action)
    client:subscribe("type/" .. thingType .. "/action/" .. action, 0)
    client:subscribe("type/" .. thingType .. "/location/" .. location .. "/action/" .. action, 0)
    client:subscribe("macaddress/" .. wifi.ap.getmac() .. "/action/" .. action, 0)
    print("Subscribed to " .. action .. " action")
end

function doSubscribeGetter(client, location, getter)
    client:subscribe("getter/" .. getter, 0)
    client:subscribe("macaddress/" .. wifi.ap.getmac() .. "/getter/" .. getter, 0)
    client:subscribe("location/" .. location .. "/getter/" .. getter, 0)
    print("Subscribed to " .. getter .. " getter")
end

function register(client, name, thingType, location, actions, getters)
    actionsJSON = "[]"
    gettersJSON = "[]"
    if table.getn(actions) > 0 then
        actionsJSON = cjson.encode(actions)
    end
    if table.getn(getters) > 0 then
        gettersJSON = cjson.encode(getters)
    end
    message = 
    '{'..
    '\"name\":\"' .. name .. '\",'..
    '\"type\":\"' .. thingType .. '\",'..   
    '\"location\":\"' .. location .. '\",'..
    '\"macaddress\":\"' .. wifi.ap.getmac() .. '\",'..
    '\"actions\":' .. actionsJSON .. ','..
    '\"getters\":' .. gettersJSON ..
    '}'
    for k, action in pairs(actions) do
        doSubscribeAction(client, thingType, location, action.name)
    end

    for k, getter in pairs(getters) do
        doSubscribeGetter(client, location, getter.name)
    end
    client:publish("register", message, 0, 0, function(client) print("Registered") heartbeat(client) end)
end

function isAction(topic, action)
    return string.match(topic, "action/" .. action .. "$") ~= nil
end

function heartbeat(client)
    client:publish("heartbeat", "{\"macaddress\":\"" .. wifi.ap.getmac() .. "\"}", 0, 0, function() end)
    tmr.alarm(1, 15000, 1, function()
        client:publish("heartbeat", "{\"macaddress\":\"" .. wifi.ap.getmac() .. "\"}", 0, 0, function() end)
    end)
end