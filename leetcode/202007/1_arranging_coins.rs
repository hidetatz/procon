impl Solution {
    pub fn arrange_coins(n: i32) -> i32 {
        let mut cn = n;
        let mut line = 1;
        
        loop {
            if (cn - line < 0) {
                break;    
            }
            cn -= line;
            line += 1;
        }
        return line - 1;
    }
}
