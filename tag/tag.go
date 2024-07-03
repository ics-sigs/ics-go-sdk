package tag

import (
    "context"
    "github.com/ics-sigs/ics-go-sdk/client/methods"

    "github.com/ics-sigs/ics-go-sdk/client/types"
)

// ListAttachedTags fetches the array of tag IDs attached to the given object.
func (t *TagsService) ListAttachedTags(ctx context.Context, targetType string, ref string) ([]string, error) {
    tags := make([]string, 0)
    treeItems, err := methods.ListAttachedTags(ctx, t.RestAPITripper, targetType, ref)
    if err == nil && len(treeItems) == 1 {
        for _, v := range treeItems[0].Children {
            if v.Checked {
                tags = append(tags, v.ID)
            }
        }
    }
    return tags, err
}

func (t *TagsService) GetTag(ctx context.Context, id string) (*types.Tag, error) {
    tag, err := methods.GetTagById(ctx, t.RestAPITripper, id)
    return tag, err
}