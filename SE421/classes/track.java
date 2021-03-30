package SE421.classes;

public class track {
    public static void main(String[] args) {
        Observer1 no =  new Observer1();
        Observer2 c =  new Observer2();
        a yes = new a();
    }
}

class a{
    private String name;
    private Listener[] listeners = new Listener[10]    ;
    private int counter;


    public void setName(String s ){
        name = s;
        for (int i = 0; i < listeners.length; i++){
            if (listeners[i] != null) listeners[i].modified(s);
        }
    }
    public String getName(){
        return name;
    }
    public void attach(Listener l){
        listeners[counter++] = l;        
    }
}
class Observer1 implements Listener{
    public int value;
    public void modified(String newname) {
        System.out.println(newname);
    }   
}
class Observer2{
    public int value;
     public void modified(String newname) {
        System.out.println(newname);
    }   

}
interface Listener {
    public void modified(String newname);    
}