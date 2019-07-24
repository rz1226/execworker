<?php

$all = "";
for($i=0;$i<80;$i++){
	$all .= "sdfsdfffffffffffffffffffffffffffffffffffffffffffffffff";
}
$count = rand(5000,10000);

for($i=0;$i<$count;$i++){
	var_dump( $i);
	usleep(1000);

	if ($i%100 == 0 ){
		echo $xxx;
	}
}
