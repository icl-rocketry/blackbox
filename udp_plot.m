clear
g = 10;
% profile clear
% profile on -historysize 5000000000

[~, localIP] = system(['netsh interface ip show address | ' ...
'findstr "IP Address" | findstr "192"']);

localIP = strtrim(erase(localIP, "IP Address: "));

u = udpport("byte", "LocalHost", localIP, "LocalPort", 20001);
log = zeros(1, 20);
datalen = 20;


% while true
i = 0;
u.NumBytesAvailable
% tic
while true
    if u.NumBytesAvailable >= 20
%         u.NumBytesAvailable
        rawdata = read(u, 20, "single");
%         fprintf('Data received\n');
        framed = separate(rawdata, datalen);
        log = [log; framed];
        i = i+1;
    end
end
% toc
% profile viewer

function new = separate(raw, cols)
    rows = length(raw)/cols;
%     fprintf('%i rows, %i cols', rows, cols)
    new = zeros(rows, cols);
    
    for i = 1:1:rows
        for j = 1:1:cols
            new(i, j) = raw(((i-1)*cols)+j);
        end        
    end    
end