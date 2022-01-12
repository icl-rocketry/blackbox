clear
[~, localIP] = system(['netsh interface ip show address | ' ...
'findstr "IP Address" | findstr "192"']);

localIP = strtrim(erase(localIP, "IP Address: "))

u = udpport("byte", "LocalHost", localIP, "LocalPort", 20001);
log = []

while true
    if u.NumBytesAvailable > 0
        fprintf('Data received\n');
        rawdata = read(u, u.NumBytesAvailable, "single");
        log = [log, reshape(rawdata, [], 19)];
    end
end
