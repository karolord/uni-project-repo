import java.awt.Color;
import java.awt.FlowLayout;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JMenu;
import javax.swing.JMenuBar;
import javax.swing.JMenuItem;


public class App extends JFrame implements ActionListener {
    JButton[] btns = new JButton[5];
    int[] arr = new int[10];
    int[] arr2 = {1,1,0,1,1,0,1,1,1,1};
    JMenuBar menuBar;
    JMenu menu;
    JMenuItem item;
    JMenuItem item2;
    public App(){
        menuBar = new JMenuBar();
        menu = new JMenu("filesystem");
        item = new JMenuItem("alpha");
        item2 = new JMenuItem("Omega");
        item2.addActionListener(
            new ActionListener(){
                @Override
                public void actionPerformed(ActionEvent e){
                    for (int i = 0; i < btns.length; i++) {
                        if(arr2[i] == 1){
                            btns[i].setBackground(Color.RED);
                        }else{
                            btns[i].setBackground(Color.GREEN);
                        }
                    }
                }
            }
        );

        item.addActionListener(
            new ActionListener(){
                @Override
                public void actionPerformed(ActionEvent e){
                    for (int i = 0; i < btns.length; i++) {
                        if(arr[i] == 1){
                            btns[i].setBackground(Color.RED);
                        }else{
                            btns[i].setBackground(Color.GREEN);
                        }
                    }
                }
            }
        );
        menuBar.add(menu);
        menu.add(item);
        menu.add(item2);
        setJMenuBar(menuBar);
        setLayout(new FlowLayout());
        for (int i = 0; i < btns.length; i++) {
            btns[i] = new JButton(Integer.toString(i));
            btns[i].addActionListener(this);
            btns[i].setBackground(Color.GREEN);
            add(btns[i]);
        }

    }
    public static void main(String[] args) throws Exception {
        App window = new App();
        window.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        window.setSize(500,500);
        window.setVisible(true);
    }
    @Override
    public void actionPerformed(ActionEvent e) {
        for (int i = 0; i < btns.length; i++) {
            if(e.getSource() == btns[i]){
                if(arr[i] == 0){
                    btns[i].setBackground(Color.RED);
                    arr[i] = 1;
                }else{
                    btns[i].setBackground(Color.GREEN);
                    arr[i] = 0;
                }
                break;
            }
        }
        
    }
}