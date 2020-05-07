package main

const (
	reservedRangeSize uint32 = 10
)

var (
	lastInteger          uint32 = 0
	reservedRangeMaximum uint32 = 0
)

func generateID() uint32 {
	return 1234567
}

/*

db {
	id string
	last_allocated uint32
}

https://code2flow.com/app

GenerateID;
reserve_check:
if(local reserved range has space?) {
  allocate an int in reserved range;
  return ID
} else {
  while(reservation failes){
    read last allocated id from DB;
    increment last allocated id by new range size;
    if (db still contains old last allocated id) {
      update db with new last allocated value ;
      update local reserved range;
      goto reserve_check;
    }
  }
}
*/
