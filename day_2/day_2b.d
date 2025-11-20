import std.stdio;
import std.conv;
import std.file;
import std.array;
import std.math.algebraic;
import std.datetime.stopwatch;

bool isSafe(int[]levels) {
    int asc = ((levels[0]-levels[1])<0)? 1 : -1 ;
    bool safe=true;
    for(int i=1;i<levels.length;i++) {
        int asc2 = ((levels[i-1]-levels[i])<0)? 1 : -1 ;
        if(asc != asc2) {
            safe=false;
            break;
        }
        int diff;
        if(asc==-1) 
            diff=levels[i-1]-levels[i];
        else
            diff=levels[i]-levels[i-1];
        if((diff<1)||(diff>3)) {
            safe=false;
            break;
        }        
    }
    return safe;
}

void doit(string fname) {
    auto file = File(fname,"r");
    int count=0;
    foreach(line; file.byLineCopy()) {
        auto strLevels = line.split();
        auto levels = new int[strLevels.length];
        for(int i=0;i<strLevels.length;i++) {
            levels[i]=to!int(strLevels[i]);
        }
        if(!isSafe(levels)) {
            for(int j=0;j<levels.length;j++) {
                auto level4 = new int[0];
                for(int k=0;k<levels.length;k++) {
                    if(j!=k) {
                        level4 ~= levels[k];
                    }
                }
                if(isSafe(level4)) {
                    count++;
                    break;
                }
            }
        }
        else {
            count++;
        }
    }
    writeln(count);
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
