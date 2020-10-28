<?php
if($handle = opendir('img/')){
    while(false != ($entry = readdir($handle))){
        if($entry != '.' && $entry != '..'){
            if(strchr($entry, ".jpg")==true){
                echo "<br /><div>$entry</div><br />";
                echo "<div><img src='./log/$entry' width='200' height='230'></div>";
            }
        }
    }
}