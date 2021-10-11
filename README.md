# Harpocrates

Harpocrates was the Greek god of silence and screts. In his name, we present this generic security utility wrapper used for Tunnel.Work projects.

## Frequent Used Functionalities

### Randutil

```
safeRandStr, err := GetRandomString(16, runesComplete) // Generate a random string with 16-byte length from runesComplete
```

### Password Generator

```
safePwd, err := GetNewStrongPassword(32) // Generate a strong-enough password. Shortcut for GetRandomString(32, runesComplete)
```

### Crypto

