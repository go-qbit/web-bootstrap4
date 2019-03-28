package template

import (
	"context"

	"github.com/go-qbit/rbac"
	"github.com/go-qbit/web/handler"
)

type NavItem struct {
	Caption       string
	Link          string
	IsActive      bool
	OpenInNewPage bool
	SubItems      []NavItem
}

type Nav struct {
	Top  []NavItem
	Left []NavItem
}

type pageCtxType int8

var pageCtxKey pageCtxType = 0

type PageOptions struct {
	RootHandler handler.IHandler

	DisableStdStatic bool
	Title            string
	HeadContent      string
	NavBar           PONavBar
	EndBodyContent   string
}

type PONavBar struct {
	DisableExpandSm bool
	Classes         []string
	Brand           string
	BrandImageSrc   string
	BrandLink       string
	Auth            PONavBarAuth
}

type PONavBarAuth struct {
	Enabled    bool
	CurUser    string
	LoginLabel string
	LoginHref  string
	LogoutHref string
}

func getNav(ctx context.Context) Nav {
	activeHandlers := make(map[string]bool)
	curHandler := handler.GetCurHandler(ctx)
	for curHandler != nil {
		activeHandlers[curHandler.GetFullPath()] = true
		curHandler = curHandler.Parent()
	}

	res := Nav{
		Top: getNavRecursive(ctx, OptionsFromContext(ctx).RootHandler, activeHandlers).SubItems,
	}

	for _, item := range res.Top {
		if item.IsActive && len(item.SubItems) > 1 {
			res.Left = item.SubItems
			break
		}
	}

	return res
}

func getNavRecursive(ctx context.Context, root handler.IHandler, activeHandlers map[string]bool) NavItem {
	subItems := make([]NavItem, 0, len(root.SubHandlers()))
	for _, path := range root.SubHandlers() {
		subHandler := root.SubHandler(path)
		if caption := subHandler.Caption(ctx); caption != "" {
			subHandlers := subHandler.SubHandlers()
			if rbac.HasPermission(ctx, subHandler.RequiredPermission()) &&
				(len(subHandlers) == 0 || subHandler.FirstAllowedSubHandler(ctx) != nil) {
				subItems = append(subItems, getNavRecursive(ctx, subHandler, activeHandlers))
			}
		}
	}

	fullPath := root.GetFullPath()
	return NavItem{
		Caption:  root.Caption(ctx),
		Link:     fullPath,
		IsActive: activeHandlers[fullPath],
		SubItems: subItems,
	}
}

func OptionsToContext(ctx context.Context, opts *PageOptions) context.Context {
	return context.WithValue(ctx, pageCtxKey, opts)
}

func SetPageTitle(ctx context.Context, title string) {
	OptionsFromContext(ctx).Title = title
}

func OptionsFromContext(ctx context.Context) *PageOptions {
	return ctx.Value(pageCtxKey).(*PageOptions)
}
