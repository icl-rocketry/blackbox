clear
[~, localIP] = system(['netsh interface ip show address | ' ...
'findstr "IP Address" | findstr "192"']);

localIP = strtrim(erase(localIP, "IP Address: "))

u = udpport("byte", "LocalHost", localIP, "LocalPort", 20001);

while true
% for i = 1:1:10
    if u.NumBytesAvailable > 0
        fprintf('Data received\n');
        u.NumBytesAvailable
        data = read(u, u.NumBytesAvailable, "single");
        
    end
end