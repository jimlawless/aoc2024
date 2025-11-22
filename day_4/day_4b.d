import std.stdio;
import std.file;
import std.array;
import std.datetime.stopwatch;

void doit(string fname) {
    auto file = File(fname,"r");
    int count=0;
    string [] arr;
    int[int] a_map;
   
    foreach(line; file.byLineCopy()) {
        arr~=line;
    }
    // [row][col]
    auto one_off = [ [-1,-1],[-1,1],[1,1],[1,-1] ] ;
    for(int i=0;i<arr.length;i++) {
        for(int j=0;j<arr[i].length;j++) {
            for(int k=0;k<one_off.length;k++) {
                auto offsets=one_off[k];
                if(find_string(arr,"MAS",i,j,offsets[0],offsets[1],0)) {
                    int location=((i+offsets[0])*100000)+(j+offsets[1]);
                    if(location in a_map) 
                        a_map[location]++;
                    else
                        a_map[location]=1;                    
                }
            }
        }
    }
    foreach(int val; a_map) {
        if(val==2) 
            count++;
    }
      
    writeln(count);
    file.close();
}

bool find_string(string [] a, string srch, int row, int col, int row_inc, int col_inc, int ndx) {
    if(ndx>=srch.length)
        return true;
    if( (row<0)||(col<0)) 
        return false;
    if( (row>=a.length)||(col>=a[0].length)) 
        return false;
    if(a[row][col]==srch[ndx]) {
        return find_string(a,srch,row+row_inc,col+col_inc,row_inc,col_inc,ndx+1);
    }
    return false;
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