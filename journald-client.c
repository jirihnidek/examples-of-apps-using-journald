#include <stdio.h>
#include <syslog.h>

// Compile using: gcc journald-client.c 

int main(void) {
    syslog(LOG_NOTICE, "Hello World");
    return 0;
}

/* // Alternative solution

#include <systemd/sd-journal.h>

int main(void) { 
    sd_journal_print(LOG_NOTICE, "Hello World");

    // sd_journal_send("MESSAGE=Hello World!",
    //         "MESSAGE_ID=52fb62f99e2c49d89cfbf9d6de5e3555",
    //         "PRIORITY=5",
    //         "HOME=%s", getenv("HOME"),
    //         "TERM=%s", getenv("TERM"),
    //         "PAGE_SIZE=%li", sysconf(_SC_PAGESIZE),
    //         "N_CPUS=%li", sysconf(_SC_NPROCESSORS_ONLN),
    //         NULL);

    return 0;
}
*/
