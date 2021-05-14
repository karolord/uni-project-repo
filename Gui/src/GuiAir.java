import java.awt.Color;
import java.awt.Graphics;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;
public class GuiAir extends JPanel{
    JButton[] ButtonsA = new JButton[5];
    JButton[] ButtonsB = new JButton[5];
    JButton[] ButtonsC = new JButton[5];
    JButton[] ButtonsD = new JButton[5];
    public void SetButtons(JButton[] Buttons,String s,int strtX,int strtY){
        for (int i = 0; i < Buttons.length; i++) {
            Buttons[i] = new JButton(s + (i+1));
            Buttons[i].setBackground(Color.GREEN);
            Buttons[i].setBounds(strtX+150*i, strtY, 60, 60);
            this.add(Buttons[i]);
        }
    }
    public GuiAir(){
        setLayout(null);
        SetButtons(ButtonsA,"A",250,200);
        SetButtons(ButtonsB,"B",100,200);
        SetButtons(ButtonsC,"C",100,300);
        SetButtons(ButtonsD,"D",100,400);
        
    }
    @Override
    public void paintComponent(Graphics g){
        super.paintComponent(g);
        g.drawArc(0, 150, 600, 225, 90, 180);
        g.drawLine(275, 150, 1000, 150);
        g.drawLine(275, 375, 1000, 375);
        g.drawLine(400, 150, 550, 50);
        g.drawLine(650, 150, 700, 50);
        g.drawLine(400, 375, 550, 475);
        g.drawLine(650, 375, 700, 475);
    }
    public static void main(String[] args) {
        JFrame frame = new JFrame();
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        GuiAir arc = new GuiAir();
        frame.add(arc);
        frame.setSize(1000,600);
        frame.setVisible(true);
    }
    
}

