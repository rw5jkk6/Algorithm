fn main(){
    let mut arr = vec![4, 1, 3, 2];
    for _ in 0..arr.len(){
        for j in 0..arr.len() - 1{
            if arr[j] > arr[j+1]{
                let tmp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = tmp;
            }
        }
    }
    println!("{:?}", arr)
}
