import std.stdio;
import std.conv;
import std.file;
import std.array;
import std.algorithm;
import std.datetime.stopwatch;

void doit(string fname) {
    auto file = File(fname,"r");
    auto left = new int[0];
    auto right = new int[0];
    foreach(line; file.byLineCopy()) {
        auto words = line.split();
        left ~= to!int(words[0]);
        right ~= to!int(words[1]);
    }
    left.sort();
    right.sort();
    int sum;
    for(int i=0; i<left.length;i++) {
        if(left[i]>right[i]) {
            sum+=left[i]-right[i];
        }
        else {
            sum+=right[i]-left[i];
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
