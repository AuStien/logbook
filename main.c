#define _GNU_SOURCE
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <errno.h>
#include <string.h>
#include <time.h>
#include <fcntl.h>

int upsertDir(char *path);
int upsertFile(char *path, struct tm *time);
void printErr(int err);

const char *DAYS[] = {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"};

int main() {
  char *path = getenv("LOGBOOK_HOME");
  if (path == NULL) {
    fprintf(stderr, "missing envvar LOGBOOK_HOME\n");
    return 1;
  }

  char *editor = getenv("EDITOR");
  if (editor == NULL) {
    editor = "vim";
  }

  time_t timeNow = time(0);
  struct tm *timeLocal = localtime(&timeNow);

  char *yearPath;
  int b = asprintf(&yearPath, "%s/%d", path, timeLocal->tm_year + 1900);
  if (b == -1) {
    return 1;
  }
  int err = upsertDir(yearPath);
  if (err != 0) {
    printErr(err);
    return 1;
  }

  char *monthPath;
  b = asprintf(&monthPath, "%s/%d", yearPath, timeLocal->tm_mon + 1);
  free(yearPath);
  if (b == -1) {
    return 1;
  }

  err = upsertDir(monthPath);
  if (err != 0) {
    free(monthPath);
    printErr(err);
    return 1;
  }

  char *filePath;
  b = asprintf(&filePath, "%s/%d.md", monthPath, timeLocal->tm_mday);
  free(monthPath);
  if (b == -1) {
    return 1;
  }

  int fd = upsertFile(filePath, timeLocal);
  if (fd == -1) {
    free(filePath);
    printErr(errno);
    return 1;
  }

  dprintf(fd, "\n## %02d:%02d\n\n", timeLocal->tm_hour, timeLocal->tm_min);

  close(fd);

  char *command;
  b = asprintf(&command, "%s + %s", editor, filePath);
  free(filePath);
  if (b == -1) {
    return 1;
  }

  int status = system(command);
  free(command);
  if (status == -1) {
    printErr(errno);
    return 1;
  } else if (status != 0) {
    printErr(status);
    return status;
  }

  return 0;
}

int upsertDir(char *path) {
  struct stat pathStat;
  int err = stat(path, &pathStat);
  if (err == -1) {
    if (errno != ENOENT) {
      return errno;
    }

    err = mkdir(path, 0775);
    if (err == -1) {
      return errno;
    }
  } else {
    if (!S_ISDIR(pathStat.st_mode)) {
      fprintf(stderr, "%s is not a directory\n", path);   
      return ENOTDIR;
    }
  }

  return 0;
}

// upsertFile makes sure a file exists.
// If it was created a header with DD/MM/YY
// will be created.
// Returns a file descriptor, or -1 on error.
// errno will be populated?
int upsertFile(char *path, struct tm *time) {
  int wasCreated = 0;

  struct stat pathStat;
  int err = stat(path, &pathStat);
  if (err == -1) {
    if (errno != ENOENT) {
      return errno;
    }
    wasCreated = 1;
  }

  int fd = open(path, O_CREAT|O_WRONLY|O_APPEND, 0775);
  if (fd == -1) {
    return -1;
  }

  if (wasCreated == 1) {
    dprintf(fd, "# %s %02d/%02d/%d\n", DAYS[time->tm_wday], time->tm_mday, time->tm_mon+1, time->tm_year % 100);
  }

  return fd;
}

void printErr(int err) {
  printf("err: %d %s\n", err, strerror(err));
}


