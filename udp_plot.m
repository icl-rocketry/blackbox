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

lw = 2;
col1 = '#FF9000';
col2 = '#2294A4';
col3 = '#C14953';

figure
y1 = animatedline('MaximumNumPoints',100, 'Color', col1,'LineWidth', lw, 'DisplayName', 'X');
y2 = animatedline('MaximumNumPoints',100, 'Color', col2,'LineWidth', lw, 'DisplayName', 'Y');
y3 = animatedline('MaximumNumPoints',100, 'Color', col3,'LineWidth', lw, 'DisplayName', 'Z');
legend('Location', 'northwest')
% ax = axis([0, 15, -20, 20]);

title('Android Accelerometer Data');
xlabel('Time (s)');
ylabel('Acceleration (m/s^2)');


% while true
i = 0;
u.NumBytesAvailable
% tic
while true
    if u.NumBytesAvailable >= 20
%         u.NumBytesAvailable
        rawdata = read(u, 20, "single");
%         fprintf('Data received\n');
        log = [log; rawdata];
        i = i+1;
        addpoints(y1, rawdata(1), rawdata(8));
        addpoints(y2, rawdata(1), rawdata(9));
        addpoints(y3, rawdata(1), rawdata(10));
        axis([rawdata(1)-11, rawdata(1)+1, -4*g, 4*g])
        drawnow
%         disp('Updated');
    end
end
% toc
% profile viewer