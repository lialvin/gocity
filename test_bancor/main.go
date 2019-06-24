package main;
 
import (
    "net"
    "time"
    "log"
    "strings"
)
 
func chkError(err error) {
    if err != nil {
        log.Fatal(err);
    }
}
 

func main() {
    #初始锚定资金
    INIT_RAM_POOL = 5000000
    #容量
    RAM_ALL = 8 * 1024 * 1024 * 1024
    # Gb
    RAM_UNIT = 1024 * 1024 * 1024
    #Kb
    RAM_PICE_UNIT  = 1024
    
    
    rank_list_proportion = []
    for i in range(96):
        rank_list_proportion.append(i/100)
    rank_list_ram = []
    
    rank_list_uos_save = []
    rank_list_price = []
    rank_uos_per_ram = []
    
    
    with open('RAM_price.csv','w') as f:
        headers = ['rank_list_ram(Gb)','rank_list_uos_save(UOS)','rank_list_price(Kb/UOS)','rank_uos_per_ram(UOS/Kb)']
        row = [0,0,0,0]
        f_csv = csv.writer(f)
        f_csv.writerow(headers)
        for i in range(len(rank_list_proportion)):
            rank_list_ram.append( RAM_ALL*rank_list_proportion[i]/(RAM_UNIT) )
            rank_list_uos_save.append( INIT_RAM_POOL*rank_list_proportion[i] / (1 - rank_list_proportion[i]) )
            rank_list_price.append(RAM_ALL*(1 - rank_list_proportion[i]) / (INIT_RAM_POOL + rank_list_uos_save[i]) / RAM_PICE_UNIT)
            rank_uos_per_ram.append(1/rank_list_price[i])
            row[0] = rank_list_ram[i]
            row[1] = rank_list_uos_save[i]
            row[2] = rank_list_price[i]
            row[3] = rank_uos_per_ram[i]
            if(i%5 == 0):
                f_csv.writerow(row)
    
    print(rank_list_ram)
    print(rank_list_uos_save)
    print(rank_list_price)
    print(rank_uos_per_ram)
    
    import matplotlib.pyplot as plt#约定俗成的写法plt
    plt.subplot(1,1,1)
    plt.plot(rank_list_ram,rank_uos_per_ram,"b", alpha=0.7)
    plt.xlabel('ram(Gb)')
    plt.ylabel('price(UOS/Kb)')
    
    # plt.subplot(1,2,2)
    # plt.plot(rank_list_ram,rank_list_uos_save,"y", alpha=0.7)
    # plt.xlabel('ram(Gb)')
    # plt.ylabel('ram_uos_pool(UOS)')
    
    plt.show()

 
}