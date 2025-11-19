import std.stdio;
import std.conv;
import std.file;
import std.array;
import std.algorithm;
import std.datetime.stopwatch;

void doit(string fname) {
    auto file = File(fname,"r");
    auto left = new int[0];
    int right = 0;
    int[int] have;
    foreach(line; file.byLineCopy()) {
        auto words = line.split();
        left ~= to!int(words[0]);
        right = to!int(words[1]);
        if(right in have) {
            have[right]=have[right]+1;
        }
        else {
            have[right]=1;
        }
    }
    int sum=0;
    for(int i=0; i<left.length;i++) {
        auto l = left[i];
        if(l in have) {
            sum+=l*have[l];
        }
    }
    writeln(sum);
    file.close();
}


void main(string[] args) {
    if(args.length < 2) {
        writeln("Syntax:\n\t",args[0]," filename\n");
        return;
    }
    auto sw = StopWatch(AutoStart.no);
    sw.start();
    try {
        doit(args[1]);
    }
    catch(FileException e) {
        writeln("Exception: ", e.msg);
    }
    sw.stop();
    writeln("Elapsed time: ",sw.peek());
}
