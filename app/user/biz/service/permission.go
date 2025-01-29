package service

import (
	"sync"

	"github.com/casbin/casbin/v2"
)

// PermissionService 权限管理服务
type PermissionService struct {
	enforcer *casbin.Enforcer
	mu       sync.RWMutex
}

var (
	permissionService *PermissionService
	once              sync.Once
)

// NewPermissionService 创建权限服务单例
func NewPermissionService(enforcer *casbin.Enforcer) *PermissionService {
	once.Do(func() {
		permissionService = &PermissionService{
			enforcer: enforcer,
		}
	})
	return permissionService
}

// AddPolicy 添加权限策略
func (s *PermissionService) AddPolicy(sub, obj, act string, eft string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.enforcer.AddPolicy(sub, obj, act, eft)
	if err != nil {
		return err
	}
	return s.enforcer.SavePolicy()
}

// RemovePolicy 删除权限策略
func (s *PermissionService) RemovePolicy(sub, obj, act string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.enforcer.RemovePolicy(sub, obj, act)
	if err != nil {
		return err
	}
	return s.enforcer.SavePolicy()
}

// AddToBlacklist 将用户添加到黑名单
func (s *PermissionService) AddToBlacklist(userID string) error {
	return s.AddPolicy(userID, ".*", ".*", "deny")
}

// RemoveFromBlacklist 将用户从黑名单移除
func (s *PermissionService) RemoveFromBlacklist(userID string) error {
	return s.RemovePolicy(userID, ".*", ".*")
}

// IsBlacklisted 检查用户是否在黑名单中
func (s *PermissionService) IsBlacklisted(userID string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.enforcer.HasPolicy(userID, ".*", ".*", "deny")
}

// ReloadPolicy 重新加载权限策略
func (s *PermissionService) ReloadPolicy() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.enforcer.LoadPolicy()
}
