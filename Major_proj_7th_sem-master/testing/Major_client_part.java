
//author :  anil.khadwal@gmail.com
//package major_client_part;

//PROJECT'S  JAVA  CODE TO RECEIVE IMAGE
import java.io.*;
import java.net.Socket;
import java.util.Calendar;
import java.util.Date;

public class Major_client_part {

    int BUFFERSIZE = 1024;
    public static void main(String[] args) {
        if(args.length<2){
            System.out.println("Enter Socket details");
            System.exit(0);
        }
        
        int port = Integer.parseInt(args[1]);
        String server = args[0];
//        BufferedOutputStream bos=null;
        Socket sock = null;
        
        try {
              System.out.println(server  +" "+  port);
              sock = new Socket(server,port);
//              InputStream in = sock.getInputStream();
//              BufferedInputStream bis = new BufferedInputStream(in);
              
         BufferedInputStream in = new BufferedInputStream(sock.getInputStream(), 1024);
         BufferedOutputStream out = new BufferedOutputStream(sock.getOutputStream(), 1024);
//         Date date = new Date();
         
         File f = new File("./received/"+server+"_"+(new Date()).toString()+".png");
         System.out.println("Receiving...");
         FileOutputStream fout = new FileOutputStream(f);
         byte[] b = new byte[1024];
         while (in.read(b) != -1) {
           fout.write(b);
//         System.out.println(b.toString()); 
         }              
            System.out.println("Done receiving anil");   
         
        } catch (Exception e) {
            System.out.println("something wrong while receiving ");
        }finally{
          try{
//               if(bos!=null) bos.close();
               if(sock!=null) sock.close();
             }catch(Exception e){
                  }
          
        }
    }
    
//    static String do_trimming(String data){
//       int i = 0;
//       String result="";
//       char s = ':';
//       while(i < data.length()){
//            if(data.charAt(i)!=s)
//                result=result+data.charAt(i);
//           i++; 
//       }
//       return result;    
//    }
    
}
