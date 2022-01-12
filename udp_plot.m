clear
[~, localIP] = system(['netsh interface ip show address | ' ...
'findstr "IP Address" | findstr "192"']);

localIP = strtrim(erase(localIP, "IP Address: "));

u = udpport("byte", "LocalHost", localIP, "LocalPort", 20001);
log = [];
datalen = 20;

while true
    if u.NumBytesAvailable >= 320
        rawdata = read(u, u.NumBytesAvailable, "single");
        fprintf('Data received\n');
        log = [log; separate(rawdata, datalen)];
    end
end

function new = separate(raw, cols)
    rows = length(raw)/cols;
%     fprintf('%i rows, %i cols', rows, cols)
    new = zeros(rows, cols);
    
    parfor i = 1:1:rows
        for j = 1:1:cols
            new(i, j) = raw(((i-1)*cols)+j);
        end        
    end    
end