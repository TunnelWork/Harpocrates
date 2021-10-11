# Harpocrates

In the name of the ancient Greek god of silence and secrets, we present Harpocrates, the generic Security Utility Wrapper by Tunnel.Work.

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

