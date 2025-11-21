import std.stdio;
import std.conv;
import std.file;
import std.array;
import std.regex;
import std.conv;
import std.datetime.stopwatch;

void doit(string fname) {
    auto file = File(fname,"r");
    auto pattern = regex(r"[mM][uU][lL]\(\d*\,\d*\)");
    int sum=0;
    foreach(line; file.byLineCopy()) {
        auto matches = matchAll( line, pattern );
        foreach(match; matches ) {
            auto op = match[0];
            op = replace(op,'(',' ');
            op = replace(op,',',' ');
            op = replace(op,')',' ');
            auto tokens = op.split();
            sum += ( (to!int(tokens[1]))*(to!int(tokens[2])) );
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