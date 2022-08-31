package tivitSftp

// SftpServer is an interface to SFTP Server implementation
type SftpServer interface {
    Run() error
    Close() error
}

// SftpClient is an interface to SFTP Client implementation

type SftpClient interface {
    Run() error
}
